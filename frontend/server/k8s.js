const Client = require('kubernetes-client').Client
const Request = require('kubernetes-client/backends/request')

const NAMESPACE = 'hackerrank-clone'

const createPodManifest = (submissionID) => ({
  "apiVersion": "v1",
  "kind": "Pod",
  "metadata": {
    "generateName": "submission-"
  },
  "spec": {
    "containers": [
      {
        "image": "patnaikshekhar/hc-clone:1.1",
        "name": "main",
        "restart": "Never",
        "env": [
          {
            "name": "SUBMISSION_ID",
            "value": `${submissionID}`,
          },
          {
            "name": "POSTGRES_PASSWORD",
            "valueFrom": {
              "secretKeyRef": {
                "name": "postgres-hc-postgresql",
                "key": "postgresql-password"
              }
            }
          },
          {
            "name": "POSTGRES_HOST",
            "value": "postgres-hc-postgresql",
          },
        ]
      }
    ],
    "restartPolicy": "Never"
  },
})

const createPod = async (submissionID) => {

  try {
    console.log('Creating pod')

    const backend = new Request(Request.config.getInCluster())
    const client = new Client({ backend, version: '1.13' })

    console.log('Creating pod 1')
    await client.api.v1.namespaces(NAMESPACE).pods.post({
      body: createPodManifest(submissionID)
    })
    console.log(await client.api.v1.namespaces(NAMESPACE).pods.get())
    console.log('Creating pod 2')
  } catch (e) {
    console.error('Create pod error', e)
  }

}

module.exports = {
  createPod
}