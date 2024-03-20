defmodule BlogWeb.PostLive do
  alias Blog.Posts
  alias Blog.Storage.LocalFilesystem
  use BlogWeb, :live_view

  def mount(_params, _session, socket) do
    {:ok, socket}
  end

  def handle_params(%{"id" => id}, _uri, socket) do
    IO.inspect(self(), label: "HANDLE PARAMS ID=#{id}")

    post = Posts.get_post!(id)
    subject = Path.basename(post.subject, ".md")

    md =
      LocalFilesystem.read_post(post.subject)
      |> Earmark.as_html!(%Earmark.Options{
        code_class_prefix: "lang-",
        smartypants: false,
        escape: false,
        # markdown에 <br> 태그 같은걸 적용시켜줌
        inner_html: true
      })

    socket =
      assign(socket,
        markdown: md,
        subject: subject,
        page_title: subject
      )

    {:noreply, socket}
  end
end
