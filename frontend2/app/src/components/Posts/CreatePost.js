import React, { useState } from "react";
import { useAuth } from '../Auth/AuthContext';
import axios from "axios";

const CreatePost = ({ onPostCreated }) => {
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [error, setError] = useState("");
  const { authToken } = useAuth();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError("");

    try {
      const response = await axios.post(
        "http://localhost:8081/api/v1/posts", // Изменили порт на 8081
        { title, content },
        {
          headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${authToken}`
          }
        }
      );

      if (response.status === 201) {
        setTitle("");
        setContent("");
        onPostCreated?.();
      }
    } catch (err) {
      console.error("Error details:", err.response?.data || err.message);
      setError(err.response?.data?.message || "Failed to create post. Please try again.");
    }
  };

  return (
    <div className="container">
      <h2>Create Post</h2>
      {error && <div className="alert alert-danger">{error}</div>}
      <form onSubmit={handleSubmit}>
        <div className="mb-3">
          <label className="form-label">Title:</label>
          <input
            type="text"
            className="form-control"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            required
          />
        </div>
        <div className="mb-3">
          <label className="form-label">Content:</label>
          <textarea
            className="form-control"
            rows="5"
            value={content}
            onChange={(e) => setContent(e.target.value)}
            required
          />
        </div>
        <button type="submit" className="btn btn-primary">
          Create Post
        </button>
      </form>
    </div>
  );
};

export default CreatePost;