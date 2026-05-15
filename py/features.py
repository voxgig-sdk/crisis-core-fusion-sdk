# CrisisCoreFusion SDK feature factory

from feature.base_feature import CrisisCoreFusionBaseFeature
from feature.test_feature import CrisisCoreFusionTestFeature


def _make_feature(name):
    features = {
        "base": lambda: CrisisCoreFusionBaseFeature(),
        "test": lambda: CrisisCoreFusionTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
