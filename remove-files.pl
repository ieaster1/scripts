#!/usr/bin/perl

use strict;
use warnings;

my $logFile = "remove-output.log";
my @files = <*>;

for my $file (@files) {
    open STDOUT, '>>', $logFile;
    print "removing $file\n";
    unlink $file or do {
        open STDERR, '>>', $logFile;
        warn("Unable to remove $file: $!");
        close(STDERR);
        next;
    };
};
