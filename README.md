You can use Kubernetes `Volume`s to provide storage for your applications. There is support for multiple types of volumes in Kubernetes. One way of categorizing them is as follows

- **Emphemeral** - `Volume`s which are tightly coupled with the `Pod` lifetime (e.g. `emptyDir` volume) i.e. they are deleted if the `Pod` is removed (for any reason). 
- **Persistent** - `Volume`s which are meant for long term storage and independent of the `Pod` or the `Node` lifecycle.  This could be `NFS` or cloud based storage in case of managed Kubernetes offerings such as Azure Kubernetes Service, Google Kubernetes Engine etc.
 
In this blog post, we will look at an example of how to use [Azure Disk](https://azure.microsoft.com/services/storage/disks/?WT.mc_id=data-0000-abhishgu) as a storage medium for your apps deployed to [Azure Kubernetes Service](https://azure.microsoft.com/services/kubernetes-service/?WT.mc_id=data-0000-abhishgu).

You will:

- Setup a Kubernetes cluster on Azure
- Create an Azure Disk and a corresponding `PersistentVolume`
- Create a `PersistentVolumeClaim` for the app `Deployment`
- Test things out to see how it all works end to end