import requests
import json
import time
from datetime import datetime

# Base URL of your server
BASE_URL = "http://localhost:8080"
session = requests.Session()


def register_user():
    """Function to register a new test user"""
    print("\n=== Registering Test User ===")

    test_user = {
        "login": f"testuser_{int(time.time())}",  # Create unique username based on timestamp
        "password": "secure_password_123",
        "email": f"test_{int(time.time())}@example.com"  # Create unique email based on timestamp
    }

    url = f"{BASE_URL}/users/register"
    headers = {"Content-Type": "application/json"}

    print("registering test user: ", test_user)
    # Make the request
    response = session.post(url, json=test_user, headers=headers)

    # Print response details
    print(f"Status Code: {response.status_code}")
    print(f"Response: {response.text}")

    # Assert expected behavior
    if response.status_code == 201:
        print("✅ User registration successful!")
    else:
        print("❌ User registration failed.")

    return response.status_code == 201


def test_create_post():
    """Test create post endpoint"""
    print("\n=== Testing Create Post ===")

    post_data = {
        "title": "Initial Post",
        "description": "This is a test post",
        "creator_id": "1",
        "is_private": False,
        "tags": ["api", "test"]
    }

    url = f"{BASE_URL}/posts/create"
    headers = {"Content-Type": "application/json"}

    response = session.post(url, json=post_data, headers=headers)

    print(f"Status Code: {response.status_code}")
    print(f"Response: {response.text}")

    if response.status_code == 200:
        print("✅ Post creation successful!")
        post_id = response.json().get("post_id")
    else:
        print("❌ Post creation failed.")
        post_id = None

    return post_id

def test_list_posts():
    """Test list posts endpoint"""
    print("\n=== Testing List Posts ===")

    url = f"{BASE_URL}/posts"

    response = session.get(url)

    print(f"Status Code: {response.status_code}")
    print(f"Response: {response.text}")

    if response.status_code == 200:
        print("✅ List posts successful!")
    else:
        print("❌ List posts failed.")

    return response.status_code == 200

def test_get_specific_post(post_id):
    """Test get specific post endpoint"""
    print("\n=== Testing Get Specific Post ===")

    url = f"{BASE_URL}/posts/{post_id}"

    response = session.get(url)

    print(f"Status Code: {response.status_code}")
    print(f"Response: {response.text}")

    if response.status_code == 200:
        print("✅ Get specific post successful!")
    else:
        print("❌ Get specific post failed.")

    return response.status_code == 200

def test_update_post(post_id):
    """Test update post endpoint"""
    print("\n=== Testing Update Post ===")

    url = f"{BASE_URL}/posts/{post_id}"
    headers = {"Content-Type": "application/json"}
    update_data = {
        "title": "Updated Post Title",
        "description": "Updated post description",
        "is_private": True,
        "tags": ["updated", "api"]
    }

    response = session.put(url, json=update_data, headers=headers)

    print(f"Status Code: {response.status_code}")
    print(f"Response: {response.text}")

    if response.status_code == 200:
        print("✅ Update post successful!")
    else:
        print("❌ Update post failed.")

    return response.status_code == 200

def test_delete_post(post_id):
    """Test delete post endpoint"""
    print("\n=== Testing Delete Post ===")

    url = f"{BASE_URL}/posts/{post_id}"

    response = session.delete(url)

    print(f"Status Code: {response.status_code}")
    print(f"Response: {response.text}")

    if response.status_code == 200:
        print("✅ Delete post successful!")
    else:
        print("❌ Delete post failed.")
        return False

    response = session.get(url)

    print(f"Status Code: {response.status_code}")
    print(f"Response: {response.text}")

    if response.status_code == 404:
        print("✅ Not found post afted delete!")
    else:
        print("❌ Got wrong status code.")
        return False

    return True

def run_post_tests():
    """Run all post-related tests sequentially"""
    print("\nStarting Post API Tests...")

    # First ensure the user is registered and logged in
    if not register_user():
        print("❌ Cannot proceed with tests, user registration failed.")
        return

    # Test create post endpoint
    post_id = test_create_post()

    success = False

    # Only proceed if post creation was successful
    if post_id:
        # Test list posts endpoint
        list_posts_success = test_list_posts()

        # Test get specific post endpoint
        get_post_success = test_get_specific_post(post_id)

        # Test update post endpoint
        update_post_success = test_update_post(post_id)

        # Test delete post endpoint
        delete_post_success = test_delete_post(post_id)

        # Check if all post tests passed
        success = all([
            list_posts_success,
            get_post_success,
            update_post_success,
            delete_post_success
        ])

    # Final summary
    print("\n=== Post Test Summary ===")
    print(f"Create Post: {'✅ Passed' if post_id else '❌ Failed'}")
    print(f"List Posts: {'✅ Passed' if list_posts_success else '❌ Failed'}")
    print(f"Get Specific Post: {'✅ Passed' if get_post_success else '❌ Failed'}")
    print(f"Update Post: {'✅ Passed' if update_post_success else '❌ Failed'}")
    print(f"Delete Post: {'✅ Passed' if delete_post_success else '❌ Failed'}")

    if success:
        print("\n✅ All post-related tests passed!")
    else:
        print("\n❌ Some post-related tests failed.")

# test private access

def create_session():
    """Create a new session for each user"""
    return requests.Session()

def register_custom(session, login, email):
    """Function to register a new test user"""
    user_data = {
        "login": login,
        "password": "secure_password_123",
        "email": email
    }
    response = session.post(f"{BASE_URL}/users/register", json=user_data)
    return response.status_code == 201

def create_private_post(session):
    """Create a private post"""
    post_data = {
        "title": "Private Post",
        "description": "This is a private post",
        "is_private": True,
        "tags": ["private", "test"]
    }
    response = session.post(f"{BASE_URL}/posts/create", json=post_data)
    if response.status_code == 200:
        return response.json().get("post_id")
    return None

def test_private_post_access_restriction():
    """Test the restriction of access to a private post by another user"""
    print("\n=== Testing Private Post Access Restriction ===")
    
    # Create first user's session, register, login, and create a private post
    session_user1 = create_session()
    user1_login = f"user1_{int(time.time())}"
    user1_email = f"user1_{int(time.time())}@example.com"
    
    assert register_custom(session_user1, user1_login, user1_email), "Registration for user1 failed"
    private_post_id = create_private_post(session_user1)
    assert private_post_id is not None, "Private post creation failed"

    # Create second user's session, register, login
    session_user2 = create_session()
    user2_login = f"user2_{int(time.time())}"
    user2_email = f"user2_{int(time.time())}@example.com"
    
    assert register_custom(session_user2, user2_login, user2_email), "Registration for user2 failed"

    # Second user tries to access the private post of the first user
    response = session_user2.get(f"{BASE_URL}/posts/{private_post_id}")
    print(f"Status Code: {response.status_code}")
    print(f"Response: {response.text}")

    # Assert that accessing the private post results in a 403 Forbidden status
    assert response.status_code == 403, "Expected forbidden access, but got a different response"

    if response.status_code == 403:
        print("✅ Access restriction to the private post is successfully enforced!")
    else:
        print("❌ Access restriction to the private post failed.")

def test_view_post(post_id):
    """Test registering a post view"""
    print("\n=== Testing View Post ===")

    url = f"{BASE_URL}/posts/{post_id}/view"
    
    response = session.post(url)

    print(f"Status Code: {response.status_code}")
    print(f"Response: {response.text}")

    if response.status_code == 200:
        print("✅ View post successful!")
    else:
        print("❌ View post failed.")

    return response.status_code == 200

def test_like_post(post_id):
    """Test liking a post"""
    print("\n=== Testing Like Post ===")

    url = f"{BASE_URL}/posts/{post_id}/like"
    
    # Like the post
    response = session.post(url)

    print(f"Status Code: {response.status_code}")
    print(f"Response: {response.text}")

    if response.status_code == 200:
        print("✅ Like post successful!")
    else:
        print("❌ Like post failed.")
        return False
        
    return True

def test_create_comment(post_id):
    """Test creating a comment on a post"""
    print("\n=== Testing Create Comment ===")

    url = f"{BASE_URL}/posts/{post_id}/comments"
    headers = {"Content-Type": "application/json"}
    comment_data = {
        "text": "This is a test comment"
    }

    response = session.post(url, json=comment_data, headers=headers)

    print(f"Status Code: {response.status_code}")
    print(f"Response: {response.text}")

    if response.status_code == 201:
        print("✅ Create comment successful!")
        comment_id = response.json().get("comment_id")
    else:
        print("❌ Create comment failed.")
        comment_id = None

    return comment_id

def test_list_comments(post_id):
    """Test listing comments for a post"""
    print("\n=== Testing List Comments ===")

    url = f"{BASE_URL}/posts/{post_id}/comments"
    
    # Test with default pagination
    response = session.get(url)

    print(f"Status Code: {response.status_code}")
    print(f"Response: {response.text}")

    if response.status_code == 200:
        print("✅ List comments successful!")
    else:
        print("❌ List comments failed.")
        return False
        
    # Test with custom pagination
    url_with_pagination = f"{url}?page_number=1&page_size=5"
    response = session.get(url_with_pagination)
    
    print(f"Status Code (with pagination): {response.status_code}")
    print(f"Response (with pagination): {response.text}")

    if response.status_code == 200:
        print("✅ List comments with pagination successful!")
    else:
        print("❌ List comments with pagination failed.")
        return False
        
    return True

def test_private_post_interaction_restriction():
    """Test restrictions on interactions with private posts"""
    print("\n=== Testing Private Post Interaction Restrictions ===")
    
    # Create first user's session, register, and create a private post
    session_user1 = create_session()
    user1_login = f"user1_interaction_{int(time.time())}"
    user1_email = f"user1_interaction_{int(time.time())}@example.com"
    
    if not register_custom(session_user1, user1_login, user1_email):
        print("❌ Registration for user1 failed")
        return False
    
    private_post_id = create_private_post(session_user1)
    if not private_post_id:
        print("❌ Private post creation failed")
        return False

    # Create second user's session and register
    session_user2 = create_session()
    user2_login = f"user2_interaction_{int(time.time())}"
    user2_email = f"user2_interaction_{int(time.time())}@example.com"
    
    if not register_custom(session_user2, user2_login, user2_email):
        print("❌ Registration for user2 failed")
        return False

    # Second user tries to view the private post
    response = session_user2.post(f"{BASE_URL}/posts/{private_post_id}/view")
    print(f"View Status Code: {response.status_code}")
    if response.status_code == 403:
        print("✅ View restriction successfully enforced!")
    else:
        print(f"❌ Expected 403 for view, got {response.status_code}")
        return False

    # Second user tries to like the private post
    response = session_user2.post(f"{BASE_URL}/posts/{private_post_id}/like")
    print(f"Like Status Code: {response.status_code}")
    if response.status_code == 403:
        print("✅ Like restriction successfully enforced!")
    else:
        print(f"❌ Expected 403 for like, got {response.status_code}")
        return False

    # Second user tries to comment on the private post
    headers = {"Content-Type": "application/json"}
    comment_data = {"text": "This should not be allowed"}
    response = session_user2.post(
        f"{BASE_URL}/posts/{private_post_id}/comments", 
        json=comment_data, 
        headers=headers
    )
    print(f"Comment Status Code: {response.status_code}")
    if response.status_code == 403:
        print("✅ Comment restriction successfully enforced!")
    else:
        print(f"❌ Expected 403 for comment, got {response.status_code}")
        return False

    return True

def run_extended_post_tests():
    """Run tests for the new post-related endpoints"""
    print("\nStarting Extended Post API Tests...")

    # First ensure the user is registered and logged in
    if not register_user():
        print("❌ Cannot proceed with tests, user registration failed.")
        return

    # Test create post endpoint to get a post ID for further tests
    post_id = test_create_post()

    success = False

    # Only proceed if post creation was successful
    if post_id:
        # Test view post endpoint
        view_post_success = test_view_post(post_id)

        # Test like post endpoint
        like_post_success = test_like_post(post_id)

        # Test create comment endpoint
        comment_id = test_create_comment(post_id)
        create_comment_success = comment_id is not None

        # Test list comments endpoint
        list_comments_success = test_list_comments(post_id)

        # Check if all post tests passed
        success = all([
            view_post_success,
            like_post_success,
            create_comment_success,
            list_comments_success
        ])

    # Test private post interaction restrictions
    private_post_restriction_success = test_private_post_interaction_restriction()

    # Final summary
    print("\n=== Extended Post Test Summary ===")
    print(f"View Post: {'✅ Passed' if view_post_success else '❌ Failed'}")
    print(f"Like Post: {'✅ Passed' if like_post_success else '❌ Failed'}")
    print(f"Create Comment: {'✅ Passed' if create_comment_success else '❌ Failed'}")
    print(f"List Comments: {'✅ Passed' if list_comments_success else '❌ Failed'}")
    print(f"Private Post Restriction: {'✅ Passed' if private_post_restriction_success else '❌ Failed'}")

    if success and private_post_restriction_success:
        print("\n✅ All extended post-related tests passed!")
    else:
        print("\n❌ Some extended post-related tests failed.")


if __name__ == "__main__":
    # Then run post-related tests
    run_post_tests()
    test_private_post_access_restriction()
    run_extended_post_tests()
