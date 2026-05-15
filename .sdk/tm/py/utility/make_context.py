# CrisisCoreFusion SDK utility: make_context

from core.context import CrisisCoreFusionContext


def make_context_util(ctxmap, basectx):
    return CrisisCoreFusionContext(ctxmap, basectx)
