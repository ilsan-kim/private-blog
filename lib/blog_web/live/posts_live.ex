defmodule BlogWeb.PostsLive do
  use BlogWeb, :live_view

  alias Blog.Storage.LocalFilesystem

  @default_thumbnail "https://elixir-lang.org/images/logo/logo.png"

  def mount(_params, _seesion, socket) do
    posts =
      LocalFilesystem.read_posts(%{limit: 10})
      |> Enum.map(fn post ->
        %{subject: post.file_name, content: post.content, thumbnail: @default_thumbnail}
      end)

    IO.inspect(posts)
    socket = assign(socket, :posts, posts)

    {:ok, socket}
  end

  attr :post, :map, required: true

  def post_row(assigns) do
    ~H"""
    <div class="posts-post">
      <h2><%= @post.subject %></h2>
      <p>
        <%= raw(@post.content) %>
      </p>
      <img src={@post.thumbnail} alt={"#{@post.subject}_thumbnail"} class="posts-image" />
    </div>
    """
  end

  defp post_at(posts, idx) do
    Enum.at(posts, idx, %{
      subject: "No Posts",
      content: "No Posts",
      thumbnail: @default_thumbnail
    })
  end
end
