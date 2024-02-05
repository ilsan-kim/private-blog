defmodule Blog.PostsFixtures do
  @moduledoc """
  This module defines test helpers for creating
  entities via the `Blog.Posts` context.
  """

  @doc """
  Generate a post.
  """
  def post_fixture(attrs \\ %{}) do
    {:ok, post} =
      attrs
      |> Enum.into(%{
        file_path: "some file_path",
        preview: "some preview",
        subject: "some subject",
        thumbnail: "some thumbnail"
      })
      |> Blog.Posts.create_post()

    post
  end
end
