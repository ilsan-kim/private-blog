defmodule Blog.LocalFileSystemTest do
  alias Blog.Storage.LocalFilesystem
  require Logger
  use ExUnit.Case

  test "read_profile!" do
    profile = LocalFilesystem.read_profile!()
    assert profile == "test-profile!\n"
  end

  test "count_posts" do
    count = LocalFilesystem.count_posts()
    assert count == 2
  end

  test "read!" do
    post = LocalFilesystem.read_post("1.md")
    assert post == "im 1.md\n"
  end

  test "read no exist" do
    post = LocalFilesystem.read_post("no_exist.md")
    assert post == "enoent"
  end
end
