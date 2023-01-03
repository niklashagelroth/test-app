# Notes

## Enabling to push images to GCP artifact Registry

1. setup a GCP service account that has access to write (repository admin works and perhaps writer is enough?)
2. download service account key file
3. login using docker login
   % cat {{service account key json file}} | docker login -u \_json_key --password-stdin {{docker host, for example: europe-north1-docker.pkg.dev}}
4. make sure the ~/.docker/config.json looks something like this:
   `    {
  "auths": {
    "europe-north1-docker.pkg.dev": {
      "auth": "XXXXXX Some big base64 string value...."
    }
  }
}`
5. if not, if theree is a credsStore value for example. Remove the line with the credsStore value and run the login command again
6. Run `cat ~/.docker/config.json | base64` and copy the output
7. Create a secret kubernetes yaml with the value like this. And `kubectl apply -f` it.
   `apiVersion: v1
kind: Secret
metadata:
  name: docker-credentials
data:
  config.json: {{paste the value}}`
