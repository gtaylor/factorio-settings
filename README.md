# Factorio Settings for Go

This package contains settings-related Go structs that can be used with the standard lib's [encoding/json](https://golang.org/pkg/encoding/json/#MarshalIndent) package. This is useful for dynamically generating Factorio server settings files.

## Example usage

```go
    package main
    
    import "github.com/gtaylor/factorio-settings"
    
    func main() {
        fs := fsettings.NewFactorioServerSettings("Test Server", "This is a test!")
        fs.Name = "My Server"
        fs.Description = "This one!"
        println(fs.String())
    }
````

More commonly, you'll just be marshalling and using the output in some way:

```go
    fs := fsettings.NewFactorioServerSettings("Test Server", "This is a test!")
    fs.Name = "My Server"
    fs.Description = "This one!"
    j, err := json.Marshal(fs)
    if err != nil {
        // Blah
    }
    // do stuff with j here
````

You can also do similar things with map generation settings:

```go
    ms := fsettings.NewMapGenSettings()
    ms.PeacefulMode = true
    ms.AutoplaceControls.Coal.Richness = "high"
    j, err := json.Marshal(ms)
    if err != nil {
        // Blah
    }
    // do stuff with j here
```

## License

This package is licensed under the 3-Clause BSD. A copy is included in the `LICENSE` file.
