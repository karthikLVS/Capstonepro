from app import app

def test_get_users():
    client = app.test_client()
    response = client.get("/users")
    assert response.status_code == 200
    assert isinstance(response.json, list)
