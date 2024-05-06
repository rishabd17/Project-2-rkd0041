func executeCommand(name string, arg ...string) error {
    cmd := exec.Command(name, arg...)
    cmd.Stderr = os.Stderr
    cmd.Stdout = os.Stdout
    return cmd.Run()
}
