select solver, avg(solvetime), percentile_cont(0.5) within group(order by solvetime)
from tests group by solver
order by solver
