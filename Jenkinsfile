pipeline {
environment {
path = "path"
}
agent any
stage('Deploy') {
steps{
    dir("$path") {
    sh "deploy.sh"
    }
  }
}
stage('Publish Influx') {
steps{
influxDbPublisher customPrefix: '', customProjectName: '', jenkinsEnvParameterField: '', jenkinsEnvParameterTag: '', selectedTarget: 'Influxdb'
    }
  }
}