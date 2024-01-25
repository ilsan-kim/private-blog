defmodule Blog.LocalFileSystemTest do
  alias Blog.Storage.LocalFilesystem
  require Logger
  use ExUnit.Case

  test "read_profile!" do
    profile = LocalFilesystem.read_profile!()
    assert profile == "test-profile!\n"
  end

  test "count_posts" do
    b = Application.get_env(:blog, :posts_dir_path)
    IO.inspect(b)
    count = LocalFilesystem.count_posts()
    assert count == 2
  end
end
