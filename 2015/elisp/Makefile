EMACS ?= emacs
BATCH = $(EMACS) -Q --batch -L .
DAYS = $(if $(DAY),day$(DAY).el,$(wildcard day[0-9][0-9].el))

.PHONY: test run

test:
	$(BATCH) $(addprefix -l ,$(DAYS)) -f ert-run-tests-batch-and-exit

run:
	@test -n "$(DAY)" || (echo "usage: make run DAY=01" && exit 1)
	@$(BATCH) -l day$(DAY).el --eval "(aoc-run $(DAY) #'day$(DAY)-part1 #'day$(DAY)-part2)"
