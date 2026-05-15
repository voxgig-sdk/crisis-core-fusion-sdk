# ProjectName SDK exists test

import pytest
from crisiscorefusion_sdk import CrisisCoreFusionSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = CrisisCoreFusionSDK.test(None, None)
        assert testsdk is not None
