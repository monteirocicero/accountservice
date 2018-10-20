from random import randint
from locust import HttpLocust, TaskSet, task

class AccountBehavior(TaskSet):

    @task
    def get_accounts(self):
        self.client.get('/accounts/%d' % randint(10000, 10099), name='/accounts/[int]')
              

class WebsiteUser(HttpLocust):
    task_set = AccountBehavior
    min_wait = 1000
    max_wait = 3000
