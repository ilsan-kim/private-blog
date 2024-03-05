defmodule Blog.Repo.Migrations.RemoveThumbnailAndFilePathFromPosts do
  use Ecto.Migration

  def change do
    alter table(:posts) do
      remove :thumbnail
      remove :file_path
    end
  end
end
