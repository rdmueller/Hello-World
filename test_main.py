import pytest
from fastapi.testclient import TestClient
from main import app

client = TestClient(app)

def test_main_without_arguments():
    """Test root endpoint returns 'Hello, World!'"""
    response = client.get("/")
    assert response.status_code == 200
    assert response.json() == {"message": "Hello, World!"}

def test_main_with_argument():
    """Test endpoint with name returns 'Hello, {name}!'"""
    response = client.get("/bob")
    assert response.status_code == 200
    assert response.json() == {"message": "Hello, bob!"}

def test_main_with_different_name():
    """Test endpoint with different name returns 'Hello, {name}!'"""
    response = client.get("/alice")
    assert response.status_code == 200
    assert response.json() == {"message": "Hello, alice!"}