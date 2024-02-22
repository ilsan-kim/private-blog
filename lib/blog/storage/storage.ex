defmodule Blog.Storage.Storage do
  @callback read_post(name :: String.t()) :: binary()
  @callback read_posts(args :: map()) :: Enum.t()
  @callback read_profile!() :: binary()
  @callback count_posts() :: integer()
end
