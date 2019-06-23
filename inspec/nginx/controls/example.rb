# encoding: utf-8
# copyright: 2018, The Authors

title 'nginx examples'

describe package('nginx') do
  it { should be_installed }
end

describe port(80) do
  it { should be_listening }
  its('protocols') { should include 'tcp' }
  its('processes') { should cmp 'nginx' }
end

describe http('http://localhost') do
  its('status') { should cmp 200 }
end
