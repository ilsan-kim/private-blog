defmodule Blog.Posts.Post do
  use Ecto.Schema
  import Ecto.Changeset

  schema "posts" do
    field :subject, :string
    field :preview, :string
    field :thumbnail, :string
    field :file_path, :string

    timestamps(type: :utc_datetime)
  end

  @doc false
  def changeset(post, attrs) do
    post
    |> cast(attrs, [:subject, :preview, :thumbnail, :file_path])
    |> validate_required([:subject, :preview, :thumbnail, :file_path])
  end
end
