defmodule Blog.Storage.LocalFilesystem do
  require Logger
  @behaviour Blog.Storage.Storage

  @profile_file_path Application.compile_env(:blog, :profile_file_path)
  @posts_dir_path Application.compile_env(:blog, :posts_dir_path)

  @impl true
  @doc """
  Reads the content of a file specified by the `file_path`.

  ## Parameters

  - `file_path`: The path to the file that will be read.

  ## Returns

  The content of the file as a binary string.

  ## Examples

      iex> Blog.Storage.LocalFilesystem.read!("path/to/file")
      "File content as string"

  Raises `Atom.to_string(reason)` if there's an error.
  """
  defp read!(file_path) when is_binary(file_path) do
    case File.read(file_path) do
      {:ok, binary} ->
        binary
        |> sanitaze()

      {:error, reason} ->
        Atom.to_string(reason)
    end
  end

  @impl true
  @doc """
  Reads a specified number of blog post contents.

  ## Parameters

  - `options`: A map containing the `:limit` key which specifies the maximum number of posts to read.

  ## Returns

  A list of maps, each containing `:file_name` and `:content` keys representing the name and content of the post.

  ## Examples

      iex> Blog.Storage.LocalFilesystem.read_posts(%{limit: 2})
      [%{file_name: "post1", content: "Content of post1"}, %{file_name: "post2", content: "Content of post2"}]

  If the `:limit` is greater than the number of available posts, it will return as many as available.
  """
  def read_posts(%{limit: limit, offset: offset}) do
    posts = list_posts()
    len_posts = length(posts)

    posts
    |> Enum.drop(offset)
    |> Enum.take(min(limit, len_posts))
    |> Enum.map(fn file -> %{file_name: file, content: read_post(file)} end)
  end

  defp list_posts() do
    case File.ls(@posts_dir_path) do
      {:ok, file_list} ->
        file_list

      {:error, reason} ->
        Logger.warning("failed on File.ls call.. #{reason}")
        []
    end
  end

  @impl true
  @doc """
  Reads a blog post content by its name.

  ## Parameters

  - `post_name`: The name of the post file to read.

  ## Returns

  The content of the blog post as a binary string.

  ## Examples

      iex> Blog.Storage.LocalFilesystem.read_post("my_post")
      "Blog post content as string"
  """
  def read_post(post_name) do
    "#{@posts_dir_path}#{post_name}"
    |> read!
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

  defp sanitaze(string) do
    String.replace(string, "\t", "    ")
  end
end
