defmodule Blog.Filesystem.Filesystem do
  @callback read!(name :: String.t()) :: binary()
end
