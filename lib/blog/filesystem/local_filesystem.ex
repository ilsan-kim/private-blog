defmodule Blog.Filesystem.LocalFilesystem do
  @behaviour Blog.Filesystem.Filesystem

  @base_dir "md"

  @impl true
  def read!(file_path) when is_binary(file_path) do
    file_path = "#{@base_dir}/#{file_path}"

    case File.read(file_path) do
      {:ok, binary} -> binary
      {:error, reason} -> Atom.to_string(reason)
    end
  end
end
