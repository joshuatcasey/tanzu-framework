// Angular imports
import { Component, OnInit } from '@angular/core';
import { Validators } from '@angular/forms';

// Library imports
import * as sortPaths from 'sort-paths';

// App imports
import AppServices from '../../../../shared/service/appServices';
import { StepFormDirective } from '../../wizard/shared/step-form/step-form';
import { TanzuEvent, TanzuEventType } from '../../../../shared/service/Messenger';
import { ValidationService } from '../../wizard/shared/validation/validation.service';
import { VSphereDatastore } from '../../../../swagger/models/v-sphere-datastore.model';
import { VsphereField } from "../vsphere-wizard.constants";
import { VSphereFolder } from '../../../../swagger/models/v-sphere-folder.model';
import { VSphereResourcePool } from '../../../../swagger/models';
import { VsphereResourceStepMapping } from './resource-step.fieldmapping';

export interface ResourcePool {
    moid?: string;
    name?: string;
    checked?: boolean;
    disabled?: boolean;
    path: string;
    parentMoid: string;
    label?: string;
    resourceType: string;
    isExpanded: boolean;
    children?: Array<ResourcePool>;
  }

enum ResourceType {
    CLUSTER = 'cluster',
    DATACENTER = 'datacenter',
    HOST = 'host',
};

@Component({
    selector: 'app-resource-step',
    templateUrl: './resource-step.component.html',
    styleUrls: ['./resource-step.component.scss']
})
export class ResourceStepComponent extends StepFormDirective implements OnInit {

    loadingResources: boolean = false;
    resourcesFetch: number = 0;
    resourcePools: Array<ResourcePool> = [];
    computeResources: Array<any> = [];
    datastores: Array<VSphereDatastore> = [];
    vmFolders: Array<VSphereFolder> = [];

    treeData = [];
    private currentDataCenter: string;

    constructor(private validationService: ValidationService) {
        super();
    }

    ngOnInit() {
        super.ngOnInit();
        AppServices.userDataFormService.buildForm(this.formGroup, this.wizardName, this.formName, VsphereResourceStepMapping);
        this.htmlFieldLabels = AppServices.fieldMapUtilities.getFieldLabelMap(VsphereResourceStepMapping);
        this.storeDefaultLabels(VsphereResourceStepMapping);
        this.registerStepDescriptionTriggers({ clusterTypeDescriptor: true,
            fields: [VsphereField.RESOURCE_DATASTORE, VsphereField.RESOURCE_POOL, VsphereField.RESOURCE_VMFOLDER]});
        this.registerDefaultFileImportedHandler(this.eventFileImported, VsphereResourceStepMapping);
        this.registerDefaultFileImportErrorHandler(this.eventFileImportError);

        this.listenToEvents();
        this.subscribeToServices();
    }

    private listenToEvents() {
        AppServices.messenger.subscribe<string>(TanzuEventType.VSPHERE_DATACENTER_CHANGED, this.onDataCenterChange.bind(this),
            this.unsubscribe);
    }

    loadResourceOptions() {
        this.resourcesFetch = 0;
        this.loadingResources = true;
        this.retrieveResourcePools(this.currentDataCenter);
        this.retrieveComputeResources(this.currentDataCenter);
        this.retrieveDatastores(this.currentDataCenter);
        this.retrieveVMFolders(this.currentDataCenter);
    }

    // public only for testing
    onDataCenterChange(event: TanzuEvent<string>) {
        this.currentDataCenter = event.payload;

        // TODO: the following setField() calls should be done when the resources have finished being fetched (not now)
        // NOTE: because the saved data values MAY be applicable to the just-chosen DC,
        // we try to set the fields to the saved value
        const fieldsToReset = [VsphereField.RESOURCE_POOL, VsphereField.RESOURCE_DATASTORE, VsphereField.RESOURCE_VMFOLDER];
        fieldsToReset.forEach(field => {
            this.setFieldWithStoredValue(field, VsphereResourceStepMapping, this.getFieldValue(field),
                { onlySelf: true, emitEvent: false});
        });
    }

    // public only for testing
    retrieveResourcePools(dc: string) {
        AppServices.messenger.publish({
            type: TanzuEventType.VSPHERE_GET_RESOURCE_POOLS,
            payload: {dc}
        });
    }

    // public only for testing
    retrieveComputeResources(dc: string) {
        AppServices.messenger.publish({
            type: TanzuEventType.VSPHERE_GET_COMPUTE_RESOURCE,
            payload: {dc}
        });
    }

    // public only for testing
    retrieveDatastores(dc: string) {
        AppServices.messenger.publish({
            type: TanzuEventType.VSPHERE_GET_DATA_STORES,
            payload: {dc}
        });
    }

    // public only for testing
    retrieveVMFolders(dc: string) {
        AppServices.messenger.publish({
            type: TanzuEventType.VSPHERE_GET_VM_FOLDERS,
            payload: {dc}
        });
    }

    constructResourceTree(resources: Array<ResourcePool>): void {
        const nodeMap: Map<string, Array<ResourcePool>> = new Map(); // key is parenet moid, value is a list of child nodes.
        const resourceTree: Array<ResourcePool> = [];
        resources.forEach(resource => {
            const parentMoid = resource.parentMoid;
            if (parentMoid) {
                if (nodeMap.has(parentMoid)) {
                    nodeMap.get(parentMoid).push(resource);
                } else {
                    nodeMap.set(parentMoid, [resource]);
                }
            } else {
                resourceTree.push(resource); // it contains root nodes
            }
            resource.label = resource.name;
        });
        const selectResourcePool = this.getStoredValue(VsphereField.RESOURCE_POOL, VsphereResourceStepMapping, '');
        this.constructTree(resourceTree, nodeMap, selectResourcePool);
        this.treeData = this.removeDatacenter(resourceTree);
    }

    constructTree(treeNodes: Array<ResourcePool>, map: Map<string, Array<ResourcePool>>, selectResourcePool: string): void {
        if (!treeNodes || treeNodes.length <= 0) {
            return;
        }

        treeNodes.forEach(node => {
            if (node.resourceType === ResourceType.HOST || node.resourceType === ResourceType.CLUSTER) {
                node.path += '/Resources';
            }
            const childNodes = map.get(node.moid) || [];
            node.children = childNodes;
            node.isExpanded = true;
            node.checked = selectResourcePool === node.path;
            this.constructTree(childNodes, map, selectResourcePool);
        });
    }

    removeDatacenter(resourceTree: Array<ResourcePool>): Array<ResourcePool> {
        let rootNodes = [];
        resourceTree.forEach(resource => {
            if (resource.resourceType === ResourceType.DATACENTER) {
                if (resource.children && resource.children.length > 0) {
                    rootNodes = [...rootNodes, ...resource.children];
                }
            } else {
                rootNodes.push(resource);
            }
        });
        return rootNodes;
    }

    handleOnClick = (selected: ResourcePool) => {
        let resourcePoolValue = '';
        if (selected.checked) {
            this.ensureOnlyOneResourceSelected(this.treeData, selected);
            resourcePoolValue = selected.path;
        }
        this.formGroup.get(VsphereField.RESOURCE_POOL).setValue(resourcePoolValue);
    }

    ensureOnlyOneResourceSelected(resourcePools: Array<ResourcePool>, selected: ResourcePool) {
        if (!resourcePools) {
            return;
        }
        resourcePools.forEach(node => {
            node.checked = node.moid === selected.moid;
            this.ensureOnlyOneResourceSelected(node.children, selected);
        });
    }

    /**
     * Get the current value of VsphereField.RESOURCE_POOL
     */
    get resourcePoolValue() {
        return this.formGroup.get(VsphereField.RESOURCE_POOL).value;
    }

    /**
     * Get the current value of VsphereField.RESOURCE_VMFOLDER
     */
    get vmFolderValue() {
        return this.formGroup.get(VsphereField.RESOURCE_VMFOLDER).value;
    }

    /**
     * Get the current value of VsphereField.RESOURCE_DATASTORE
     */
    get datastoreValue() {
        return this.formGroup.get(VsphereField.RESOURCE_DATASTORE).value;
    }

    dynamicDescription(): string {
        const vmFolder = this.getFieldValue(VsphereField.RESOURCE_VMFOLDER, true);
        const datastore = this.getFieldValue(VsphereField.RESOURCE_DATASTORE, true);
        const resourcePool = this.getFieldValue(VsphereField.RESOURCE_POOL, true);
        if (vmFolder && datastore && resourcePool) {
            return 'Resource Pool: ' + resourcePool + ', VM Folder: ' + vmFolder + ', Datastore: ' + datastore;
        }
        return `Specify the resources for this ${this.clusterTypeDescriptor} cluster`;
    }

    private incrementResourcesFetched() {
        this.resourcesFetch += 1;
        if (this.resourcesFetch === 4) {
            this.loadingResources = false;
        }
    }

    private subscribeToServices() {
        AppServices.dataServiceRegistrar.stepSubscribe<VSphereResourcePool>(this, TanzuEventType.VSPHERE_GET_RESOURCE_POOLS,
            this.onFetchedResourcePools.bind(this), AppServices.dataServiceRegistrar.appendingStepErrorHandler(this) );
        AppServices.dataServiceRegistrar.stepSubscribe<ResourcePool>(this, TanzuEventType.VSPHERE_GET_COMPUTE_RESOURCE,
            this.onFetchedComputeResources.bind(this), AppServices.dataServiceRegistrar.appendingStepErrorHandler(this) );
        AppServices.dataServiceRegistrar.stepSubscribe<VSphereDatastore>(this, TanzuEventType.VSPHERE_GET_DATA_STORES,
            this.onFetchedDatastores.bind(this), AppServices.dataServiceRegistrar.appendingStepErrorHandler(this) );
        AppServices.dataServiceRegistrar.stepSubscribe<VSphereFolder>(this, TanzuEventType.VSPHERE_GET_VM_FOLDERS,
            this.onFetchedVmFolders.bind(this), AppServices.dataServiceRegistrar.appendingStepErrorHandler(this) );
    }

    private sortVsphereResources(data: any[]): any[] {
        return sortPaths(data, function (item) { return item.name; }, '/');
    }

    private onFetchedVmFolders(data: VSphereFolder[]) {
        this.incrementResourcesFetched();
        if (data === null || data === undefined) {
            data = [];
        }
        this.vmFolders = this.sortVsphereResources(data);
        const storedVmFolder = this.getStoredValue(VsphereField.RESOURCE_VMFOLDER, VsphereResourceStepMapping, '');
        const selectValue = (data.length === 1) ? data[0].name : storedVmFolder;
        const validators = [Validators.required,
            this.validationService.isValidNameInList(data.map(vmFolder => vmFolder.name))];
        this.resurrectField(VsphereField.RESOURCE_VMFOLDER, validators, selectValue);
    }

    private onFetchedDatastores(data: VSphereDatastore[]) {
        this.incrementResourcesFetched();
        if (data === null || data === undefined) {
            data = [];
        }
        this.datastores = this.sortVsphereResources(data);
        const selectValue = (data.length === 1) ? data[0].name :
            this.getStoredValue(VsphereField.RESOURCE_DATASTORE, VsphereResourceStepMapping, '');
        const validators = [Validators.required,
            this.validationService.isValidNameInList(data.map(vmFolder => vmFolder.name))];
        this.resurrectField(VsphereField.RESOURCE_DATASTORE, validators, selectValue);
    }

    private onFetchedComputeResources(data: ResourcePool[]) {
        this.incrementResourcesFetched();
        if (data === null || data === undefined) {
            data = [];
        }
        this.computeResources = this.sortVsphereResources(data);
        this.constructResourceTree(data);
    }

    private onFetchedResourcePools(data: VSphereResourcePool[]) {
        this.incrementResourcesFetched();
        if (data === null || data === undefined) {
            data = [];
        }
        this.resourcePools = this.sortVsphereResources(data);
    }

    protected storeUserData() {
        super.storeUserDataFromMapping(VsphereResourceStepMapping);
        super.storeDefaultDisplayOrder(VsphereResourceStepMapping);
    }
}
