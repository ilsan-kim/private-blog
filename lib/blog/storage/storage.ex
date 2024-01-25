defmodule Blog.Storage.Storage do
  @callback read!(name :: String.t()) :: binary()
  @callback read_profile!() :: binary()
  @callback count_posts() :: integer()
end
