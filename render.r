args = commandArgs(trailingOnly=TRUE)

library(ggplot2)
library(ggthemes)
library(ggthemr)
library(scales)

theme_update(
  panel.background = element_blank(),
  panel.grid.major.y = element_line(color = "grey"),
  panel.grid = element_blank(),
  text = element_text(family = "serif")
)

d <- read.table(args[1], sep=",", header=TRUE)

d <- subset(d, status == "200")

d$status <- factor(d$status)
d$simul <- factor(d$simul)

# p <- (
#   ggplot(d, aes(duration, colour = factor(rate)))
#   + geom_histogram(binwidth = 0.01)
#   # + geom_violin()
#   # + geom_point(size = 0.1)
#   # + geom_jitter(size = 0.1)
#   # + labs(x = "Request Start (s)", y = "Request Duration (s)", colour = "Status")
#   # + scale_y_continuous(breaks = pretty_breaks(10))
#   + scale_x_continuous(breaks = pretty_breaks(20))
#   # + facet_wrap(vars(status))
# )

p <- (
  ggplot(d, aes(stop, duration, color = factor(simul)))
  # + geom_histogram(binwidth = 0.01)
  # + geom_violin()
  + geom_point(size = 0.1)
  # + geom_jitter(size = 0.1)
  # + labs(x = "Request Start (s)", y = "Request Duration (s)", colour = "Status")
  # + scale_y_continuous(breaks = pretty_breaks(10))
  # + scale_y_log10()
  # + ylim(c(0.1, 0.3))
  + ylim(c(0.55, 0.6))
  + scale_x_continuous(breaks = pretty_breaks(10))
  # + facet_wrap(vars(simul))
)

ggsave(args[2], p, width = 10, height = 4, dpi = 160)
