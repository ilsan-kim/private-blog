defmodule Blog.Repo.Migrations.CreatePosts do
  use Ecto.Migration

  def change do
    create table(:posts) do
      add :subject, :string
      add :preview, :string
      add :thumbnail, :string
      add :file_path, :string

      timestamps(type: :utc_datetime)
    end
  end
end
