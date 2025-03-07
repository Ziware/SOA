import requests
import json
import time
from datetime import datetime

# Base URL of your server
BASE_URL = "http://localhost:8080"
session = requests.Session()

# Test data
test_user = {
    "login": f"testuser_{int(time.time())}",  # Create unique username based on timestamp
    "password": "secure_password_123",
    "email": f"test_{int(time.time())}@example.com"  # Create unique email based on timestamp
}

profile_data = {
    "first_name": "John",
    "last_name": "Doe",
    "birth_date": "1990-01-01",
    "email": f"john_{int(time.time())}@example.com",
    "phone_number": "+1234567890"
}

def test_register():
    """Test user registration endpoint"""
    print("\n=== Testing User Registration ===")
    
    url = f"{BASE_URL}/users/register"
    headers = {"Content-Type": "application/json"}
    
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

def test_login():
    """Test user login endpoint"""
    print("\n=== Testing User Login ===")
    
    url = f"{BASE_URL}/users/login"
    headers = {"Content-Type": "application/json"}
    login_data = {
        "login": test_user["login"],
        "password": test_user["password"]
    }
    
    # Make the request
    response = session.post(url, json=login_data, headers=headers)
    
    # Print response details
    print(f"Status Code: {response.status_code}")
    print(f"Response: {response.text}")
    
    # Show cookies
    print(f"Cookies: {session.cookies.get_dict()}")
    
    # Assert expected behavior
    if response.status_code == 201:
        print("✅ User login successful!")
    else:
        print("❌ User login failed.")
    
    return response.status_code == 201

def test_get_profile():
    """Test get user profile endpoint"""
    print("\n=== Testing Get User Profile ===")
    
    url = f"{BASE_URL}/users/profile"
    
    # Make the request (cookies are automatically included by the session)
    response = session.get(url)
    
    # Print response details
    print(f"Status Code: {response.status_code}")
    print(f"Response: {response.text}")
    
    # Assert expected behavior
    if response.status_code == 200:
        print("✅ Get profile successful!")
    else:
        print("❌ Get profile failed.")
    
    return response.status_code == 200

def test_update_profile():
    """Test update user profile endpoint"""
    print("\n=== Testing Update User Profile ===")
    
    url = f"{BASE_URL}/users/profile"
    headers = {"Content-Type": "application/json"}
    
    # Make the request (cookies are automatically included by the session)
    response = session.put(url, json=profile_data, headers=headers)
    
    # Print response details
    print(f"Status Code: {response.status_code}")
    print(f"Response: {response.text}")
    
    # Assert expected behavior
    if response.status_code == 200:
        print("✅ Profile update successful!")
    else:
        print("❌ Profile update failed.")
    
    return response.status_code == 200

def run_all_tests():
    """Run all tests in sequence"""
    print("Starting API Tests...")
    
    # Test register endpoint
    register_success = test_register()
    
    # Test login endpoint
    login_success = test_login()
    
    # Only test profile endpoints if login was successful
    if login_success:
        # Test get profile endpoint
        get_profile_success = test_get_profile()
        
        # Test update profile endpoint
        update_profile_success = test_update_profile()
        
        # Check if all authenticated tests passed
        if get_profile_success and update_profile_success:
            print("\n✅ All authenticated tests passed!")
        else:
            print("\n❌ Some authenticated tests failed.")
    else:
        print("\n❌ Skipping profile tests due to login failure.")
    
    # Final summary
    print("\n=== Test Summary ===")
    print(f"Registration: {'✅ Passed' if register_success else '❌ Failed'}")
    print(f"Login: {'✅ Passed' if login_success else '❌ Failed'}")
    if login_success:
        print(f"Get Profile: {'✅ Passed' if get_profile_success else '❌ Failed'}")
        print(f"Update Profile: {'✅ Passed' if update_profile_success else '❌ Failed'}")

# Run all tests
if __name__ == "__main__":
    run_all_tests()
