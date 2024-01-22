defmodule BlogWeb.BlogLive do
  use BlogWeb, :live_view

  def mount(_params, _session, socket) do
    md =
      File.read!("a.md")
      |> Earmark.as_html!(%Earmark.Options{
        code_class_prefix: "lang-",
        smartypants: false,
        escape: false
      })

    socket =
      assign(socket,
        markdown: md
      )

    {:ok, socket}
  end

  def render(assigns) do
    ~H"""
    <h1>Bingo Boss 📢</h1>
    <body class="line-numbers">
      <div id="md">
        <%= raw(@markdown) %>
      </div>
    </body>
    """
  end
end
