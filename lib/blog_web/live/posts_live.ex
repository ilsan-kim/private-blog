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
    pages = pages(options, LocalFilesystem.count_posts())

    socket =
      assign(
        socket,
        posts: posts,
        options: options,
        pages: pages
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

  defp pages(options, total_count) do
    total_page_count = div(total_count + options.limit - 1, options.limit)
    current_page = div(options.offset, options.limit) + 1
    has_more_page = current_page < total_page_count

    start_page = max(current_page - 2, 1)
    end_page = min(current_page + 2, total_page_count)

    pages = pages_recursive(start_page, end_page, current_page, [])

    %{
      pages: pages,
      has_more_page: has_more_page
    }
  end

  defp pages_recursive(current_page, end_page, target_page, acc) when current_page <= end_page do
    current_page? = current_page == target_page

    pages_recursive(current_page + 1, end_page, target_page, [{current_page, current_page?} | acc])
  end

  defp pages_recursive(_current_page, _end_page, _target_page, acc) do
    Enum.reverse(acc)
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
