defmodule KV.Mixfile do
  use Mix.Project

  #Used for project configuration
  def project do
    [app: :kv,
     version: "0.0.1",
     elixir: "~> 1.1",
     build_embedded: Mix.env == :prod,
     start_permanent: Mix.env == :prod,
     deps: deps]
  end

  #Generates the application file
  def application do
    [applications: [:logger]]
  end

  #Define dependencies; not required
  defp deps do
    []
  end
end
