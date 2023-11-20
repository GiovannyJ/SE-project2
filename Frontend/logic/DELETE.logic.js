/**
 * ==================================================================
 * DELETE METHODS
 * used to delete obj from database using html text input element
 * calls to delete endpoint in API
 * ==================================================================
 */

/**
 * deletes account using accountID
*/
async function deleteAccount() {
    const accountId = document.getElementById('accountId').value;
    const url = `http://localhost:8080/account/delete/${accountId}`;
  
    try {
      const response = await fetch(url, {
        method: 'DELETE',
      });
  
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
  
      alert('Account deleted successfully!');
    } catch (error) {
      console.error('Error during DELETE request:', error.message);
    }
  }
  
  /**
   * deletes post using postID
  */
  async function deletePost() {
    const postId = document.getElementById('postId').value;
    const url = `http://localhost:8080/posts/delete/${postId}`;
  
    try {
      const response = await fetch(url, {
        method: 'DELETE',
      });
  
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
  
      alert('Post deleted successfully!');
    } catch (error) {
      console.error('Error during DELETE request:', error.message);
    }
  }
  
  /**
   * Deletes comment from database using id
  */
  async function deleteComment() {
    const commentId = document.getElementById('commentId').value;
    const url = `http://localhost:8080/posts/delete/comment/${commentId}`;
  
    try {
      const response = await fetch(url, {
        method: 'DELETE',
      });
  
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
  
      alert('Comment deleted successfully!');
    } catch (error) {
      console.error('Error during DELETE request:', error.message);
    }
  }