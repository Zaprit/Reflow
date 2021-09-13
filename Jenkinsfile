node {
    checkout scm
    // Ensure the desired Go version is installed
    def root = tool type: 'go', name: 'Go 1.17.1'

    // Export environment variables pointing to the directory where Go was installed
    withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
        sh 'go get ./...'
        sh 'rm -rf build'
        sh 'mkdir build'
        sh 'go build -o build/reflow-linux-amd64 .'
        sh 'GOOS=darwin GOARCH=arm64 go build -o build/reflow-darwin-arm64 .'
        sh 'GOOS=windows GOARCH=amd64 go build -o build/reflow-windows-amd64 .'
        archiveArtifacts artifacts: 'build/*'
    }
}
