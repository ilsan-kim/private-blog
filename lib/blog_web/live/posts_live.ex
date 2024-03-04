defmodule BlogWeb.PostsLive do
  use BlogWeb, :live_view

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
    pages = pages(options, Blog.Posts.posts_count())

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
      <h2>
        <.link patch={~p"/posts/#{@post.id}"} class="link-text">
          <%= @post.subject %>
        </.link>
      </h2>
      <p>
        <%= raw(@post.content) %>
      </p>
      <img
        src={"http://localhost:8080/static/#{@post.thumbnail}"}
        alt={"#{@post.subject}_thumbnail"}
        class="posts-image"
      />
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
      id: 0,
      subject: "No Posts",
      content: "No Posts",
      thumbnail: @default_thumbnail
    })
  end

  defp get_posts(%{limit: limit, offset: offset}) do
    Blog.Posts.list_posts(%{sort_by: :updated_at, sort_order: :desc, limit: limit, offset: offset})
    |> Enum.map(&format_post/1)
  end

  defp format_post(%{thumbnail: ""} = post) do
    %{id: post.id, subject: post.subject, content: post.preview, thumbnail: @default_thumbnail}
  end

  defp format_post(post) do
    %{id: post.id, subject: post.subject, content: post.preview, thumbnail: post.thumbnail}
  end

  defp param_to_integer(nil, default), do: default

  defp param_to_integer(param, default) do
    case Integer.parse(param) do
      {number, _} -> number
      :error -> default
    end
  end
end
