defmodule BlogWeb.PostsLive do
  use BlogWeb, :live_view

  alias Blog.Storage.LocalFilesystem

  @default_thumbnail "https://elixir-lang.org/images/logo/logo.png"

  # 어차피 mount 직후 handle_params 가 호출되기 때문에 여기서 posts를 조회할 필요가 없음
  def mount(_params, _seesion, socket) do
    {:ok, socket, temporary_assigns: [posts: []]}
  end

  def handle_params(params, _uri, socket) do
    limit = param_to_integer(params["limit"], 10)
    offset = param_to_integer(params["offset"], 0)

    options = %{limit: limit, offset: offset}
    posts = get_posts(options)

    socket =
      assign(
        socket,
        posts: posts,
        options: options
      )

    {:noreply, socket}
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

  defp get_posts(%{limit: limit, offset: offset}) do
    LocalFilesystem.read_posts(%{limit: limit, offset: offset})
    |> Enum.map(fn post ->
      %{subject: post.file_name, content: post.content, thumbnail: @default_thumbnail}
    end)
  end

  defp param_to_integer(nil, default), do: default

  defp param_to_integer(param, default) do
    case Integer.parse(param) do
      {number, _} -> number
      :error -> default
    end
  end
end
