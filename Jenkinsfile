pipeline {
    agent any
    stages {
        tools {
            go 'go-1.17.1'
        }
        stage('Compile') {
            steps {
                sh 'go build -o reflow-linux-amd64'
            }
        }
        stage('Test') {
            environment {
                CODECOV_TOKEN = credentials('codecov_token')
            }
            steps {
                sh 'go test ./... -coverprofile=coverage.txt'
                sh "curl -s https://codecov.io/bash | bash -s -"
            }
        }
        stage('Code Analysis') {
            steps {
                sh 'curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b $GOPATH/bin v1.12.5'
                sh 'golangci-lint run'
            }
        }
        stage('cleanup'){
            archiveArtifacts artifacts: ''
            rm -rf build
            discordSend description: "Reflow Jenkins Build", footer: "jenkins@sandpit", link: env.BUILD_URL, result: currentBuild.currentResult, title: JOB_NAME, webhookURL: "Webhook URL"
        }
    }

}

// node {
//     checkout scm
//     // Ensure the desired Go version is installed
//     def root = tool type: 'go', name: 'Go 1.17.1'
//
//     // Export environment variables pointing to the directory where Go was installed
//     withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
//         sh 'go get ./...'
//         sh 'rm -rf build'
//         sh 'mkdir build'
//         sh 'go build -o build/reflow-linux-amd64 .'
//         sh 'GOOS=darwin GOARCH=arm64 go build -o build/reflow-darwin-arm64 .'
//         sh 'GOOS=windows GOARCH=amd64 go build -o build/reflow-windows-amd64 .'
//         archiveArtifacts artifacts: 'build/*'
//     }
//     environment {
//                     JBOSS_CREDS = credentials('jenkins-jboss-dev-creds')
//                 }
// }
