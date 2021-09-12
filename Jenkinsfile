node {
    checkout scm
    // Ensure the desired Go version is installed
    def root = tool type: 'go', name: 'Go 1.17.1'

    // Export environment variables pointing to the directory where Go was installed
    withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
        sh 'go get ./...'
        sh 'mkdir build'
        sh 'go build -o reflow-linux-amd64 ./...'
        sh 'GOOS=darwin GOARCH=arm64 go build -o reflow-darwin-arm64 ./...'
        sh 'GOOS=win GOARCH=amd64 go build -o reflow-windows-amd64 ./...'
        archiveArtifacts artifacts: 'build/*'
    }
}
