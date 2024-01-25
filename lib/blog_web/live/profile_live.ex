defmodule BlogWeb.ProfileLive do
  use BlogWeb, :live_view

  alias Blog.Storage.LocalFilesystem

  def mount(_params, _session, socket) do
    md =
      LocalFilesystem.read_profile!()
      |> Earmark.as_html!(%Earmark.Options{
        code_class_prefix: "lang-",
        smartypants: false,
        escape: false,
        # markdown에 <br> 태그 같은걸 적용시켜줌
        inner_html: true
      })

    socket = assign(socket, markdown: md)
    {:ok, socket}
  end
end
