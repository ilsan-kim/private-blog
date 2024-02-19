defmodule BlogWeb.FileController do
  use BlogWeb, :controller

  def serve_file(conn, %{"file_path" => file_path}) do
    file = Path.join("priv/static", file_path)

    if File.exists?(file) do
      send_file(conn, 200, file)
    else
      send_resp(conn, 404, "File not found")
    end
  end
end
