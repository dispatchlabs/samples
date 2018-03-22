# Install Google Cloud CLI
- `https://cloud.google.com/sdk` - Mac/Win/Linux
- `yay google-cloud-sdk` for Arch

# Config the GC-CLI
- `gcloud init`
- if needed more docs here
	- `https://cloud.google.com/compute/docs/gcloud-compute`
	- `https://cloud.google.com/compute/docs/instances/create-start-instance`
	- `https://cloud.google.com/compute/docs/machine-types`
	- `https://cloud.google.com/sdk/gcloud/reference/compute`
- practical
	- `export CLOUDSDK_COMPUTE_ZONE=ZONE`
	- `export CLOUDSDK_COMPUTE_REGION=REGION`

# VM Manipulation
- `gcloud compute instances list`
- `gcloud compute ssh my-instance --zone us-central1-a`
- `gcloud compute scp ~/file-1 my-instance:~/remote-destination --zone us-central1-a`
- `gcloud compute scp my-instance:~/file-1 ~/local-destination --zone us-central1-a`
- `gcloud compute instances create --help`
- `gcloud compute images list | grep debian`
- `gcloud compute machine-types list | grep us-west1-a`
- Creating an instance from an image
	- `gcloud compute instances create testvm-from-gcli --image-family debian-9 --image-project debian-cloud --machine-type f1-micro`