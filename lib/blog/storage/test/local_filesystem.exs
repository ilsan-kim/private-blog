defmodule Blog.Storage.Test.LocalFilesystem do
  require Logger
  use ExUnit.Case

  test "read!" do
    a = Application.get_env(:blog, :profie_file_path)
    Logger.info(a)

    assert a == "asd"
  end
end
