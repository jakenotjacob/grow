#[KV Tutorial](http://elixir-lang.org/getting-started/mix-otp/introduction-to-mix.html)

##Mix

Elixir executable used for dependency management, packaging, preparing
documentation, testing, and more. See: `man mix`

Creating a new project with mix
===============================
To create a basic project using mix use:
`mix new projectname`
 - `--sup`    Creats a supervision tree
 - `--module` Explicitly states the **main module** name
  -  defaults to _uppercases_ the first letter, _lowercase_ remainder

Mix files
---------
####mix.exs

Project function used for project configuration
```elixir
def project do
  [app: :appname,
   #...
   deps: deps]
end
```

Application function for generating the application file
```elixir
def application do
  [applications: [:logger]]
end
```

