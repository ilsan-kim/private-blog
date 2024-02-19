defmodule Blog.Posts do
  @moduledoc """
  The Posts context.
  """

  import Ecto.Query, warn: false
  alias Blog.Repo

  alias Blog.Posts.Post

  def posts_count do
    Repo.aggregate(Post, :count, :id)
  end

  def sort(query, %{sort_by: sort_by, sort_order: sort_order}) do
    order_by(query, {^sort_order, ^sort_by})
  end

  def sort(query, _options), do: query

  def paginate(query, %{offset: offset, limit: limit}) when offset >= 0 and limit >= 0 do
    IO.inspect(offset)
    IO.inspect(limit)

    query
    |> limit(^limit)
    |> offset(^offset)
  end

  def paginate(query, _options) do
    IO.inspect("hihihi")

    query
    |> limit(12)
    |> offset(0)
  end

  @doc """
  Returns the list of posts.

  ## Examples

      iex> list_posts()
      [%Post{}, ...]

  """
  def list_posts do
    Repo.all(from p in Post, order_by: [asc: p.id])
  end

  @doc """
  Returns the list of posts with params %{sort_by: sort_by, sort_order: sort_order, offset: offset, limit: limit}

  ## Examples

      iex> list_posts(%{sort_by: "id", sort_order: :asc, offset: 0, limit: 10})
      [%Post{}, ...]
  """
  def list_posts(options) when is_map(options) do
    IO.inspect(options)

    from(Post)
    |> sort(options)
    |> paginate(options)
    |> Repo.all()
  end

  @doc """
  Gets a single post.

  Raises `Ecto.NoResultsError` if the Post does not exist.

  ## Examples

      iex> get_post!(123)
      %Post{}

      iex> get_post!(456)
      ** (Ecto.NoResultsError)

  """
  def get_post!(id), do: Repo.get!(Post, id)

  @doc """
  Creates a post.

  ## Examples

      iex> create_post(%{field: value})
      {:ok, %Post{}}

      iex> create_post(%{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def create_post(attrs \\ %{}) do
    %Post{}
    |> Post.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a post.

  ## Examples

      iex> update_post(post, %{field: new_value})
      {:ok, %Post{}}

      iex> update_post(post, %{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def update_post(%Post{} = post, attrs) do
    post
    |> Post.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a post.

  ## Examples

      iex> delete_post(post)
      {:ok, %Post{}}

      iex> delete_post(post)
      {:error, %Ecto.Changeset{}}

  """
  def delete_post(%Post{} = post) do
    Repo.delete(post)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking post changes.

  ## Examples

      iex> change_post(post)
      %Ecto.Changeset{data: %Post{}}

  """
  def change_post(%Post{} = post, attrs \\ %{}) do
    Post.changeset(post, attrs)
  end
end
