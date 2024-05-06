func EnvironmentVariables(w io.Writer, args ...string) error {
    for _, env := range os.Environ() {
        fmt.Fprintln(w, env)
    }
    return nil
}
