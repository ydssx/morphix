from locust import FastHttpUser, task,between

class HelloWorldUser(FastHttpUser):
    wait_time = between(1, 5)
    @task
    def hello_world(self):
        self.client.get("/v1/images/generate_status")
