from random import randint
from locust import HttpUser, TaskSet, task, between

class AccountBehavior(HttpUser):

    @task
    def get_accounts(self):
        random_account = randint(10000, 10099)
        self.client.get(f'/accounts/{random_account}')