defmodule Blog.PostsTest do
  use Blog.DataCase

  alias Blog.Posts

  describe "posts" do
    alias Blog.Posts.Post

    import Blog.PostsFixtures

    @invalid_attrs %{subject: nil, preview: nil, thumbnail: nil, file_path: nil}

    test "list_posts/0 returns all posts" do
      post = post_fixture()
      assert Posts.list_posts() == [post]
    end

    test "get_post!/1 returns the post with given id" do
      post = post_fixture()
      assert Posts.get_post!(post.id) == post
    end

    test "create_post/1 with valid data creates a post" do
      valid_attrs = %{subject: "some subject", preview: "some preview", thumbnail: "some thumbnail", file_path: "some file_path"}

      assert {:ok, %Post{} = post} = Posts.create_post(valid_attrs)
      assert post.subject == "some subject"
      assert post.preview == "some preview"
      assert post.thumbnail == "some thumbnail"
      assert post.file_path == "some file_path"
    end

    test "create_post/1 with invalid data returns error changeset" do
      assert {:error, %Ecto.Changeset{}} = Posts.create_post(@invalid_attrs)
    end

    test "update_post/2 with valid data updates the post" do
      post = post_fixture()
      update_attrs = %{subject: "some updated subject", preview: "some updated preview", thumbnail: "some updated thumbnail", file_path: "some updated file_path"}

      assert {:ok, %Post{} = post} = Posts.update_post(post, update_attrs)
      assert post.subject == "some updated subject"
      assert post.preview == "some updated preview"
      assert post.thumbnail == "some updated thumbnail"
      assert post.file_path == "some updated file_path"
    end

    test "update_post/2 with invalid data returns error changeset" do
      post = post_fixture()
      assert {:error, %Ecto.Changeset{}} = Posts.update_post(post, @invalid_attrs)
      assert post == Posts.get_post!(post.id)
    end

    test "delete_post/1 deletes the post" do
      post = post_fixture()
      assert {:ok, %Post{}} = Posts.delete_post(post)
      assert_raise Ecto.NoResultsError, fn -> Posts.get_post!(post.id) end
    end

    test "change_post/1 returns a post changeset" do
      post = post_fixture()
      assert %Ecto.Changeset{} = Posts.change_post(post)
    end
  end
end
