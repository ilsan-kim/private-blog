defmodule Blog.Storage.LocalFilesystem do
  require Logger
  @behaviour Blog.Storage.Storage

  @profile_file_path Application.compile_env(:blog, :profile_file_path)
  @posts_dir_path Application.compile_env(:blog, :posts_dir_path)

  @impl true
  def read!(file_path) when is_binary(file_path) do
    case File.read(file_path) do
      {:ok, binary} -> binary
      {:error, reason} -> Atom.to_string(reason)
    end
  end

  @impl true
  @doc """
  Reads the profile data from the configured profile file path.

  ## Examples

      iex> Blog.Storage.LocalFilesystem.read_profile!()
      "Profile content as string"

  """
  def read_profile!() do
    read!(@profile_file_path)
  end

  @impl true
  @doc """
  Counts the number of post files in the configured posts directory.

  ## Examples

      iex> Blog.Storage.LocalFilesystem.count_posts()
      3

  Returns the count of files or 0 if there's an error accessing the directory.
  """
  def count_posts() do
    case File.ls(@posts_dir_path) do
      {:ok, files} ->
        length(files)

      {:error, reason} ->
        Logger.warning("failed on File.ls call.. #{reason}")
        0
    end
  end
end
