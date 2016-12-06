package fs

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"path"

	"cloud.google.com/go/storage"

	. "gopkg.in/check.v1"
)

var _ = Suite(&GoogleCloudStorageSuite{})

var (
	gcsConfigFixture = "ewogICJ0eXBlIjogInNlcnZpY2VfYWNjb3VudCIsCiAgInByb2plY3RfaWQiOiAic3JjZC10ZXN0IiwKICAicHJpdmF0ZV9rZXlfaWQiOiAiZDZmMTI3MDA2ZmI5ZjY2ZTE2YjIxMzgxY2UxZGI5MDA5ODkxNGE4NCIsCiAgInByaXZhdGVfa2V5IjogIi0tLS0tQkVHSU4gUFJJVkFURSBLRVktLS0tLVxuTUlJRXZBSUJBREFOQmdrcWhraUc5dzBCQVFFRkFBU0NCS1l3Z2dTaUFnRUFBb0lCQVFDd2JzLytpajhHTUN0VFxuZ0ZZcjl4cDlmV0hnT0VVTHhrZHB3R0t6aFJpaFhLc2RndWc2d3VsSDFDOXphRFY3VlpyWWxwdjVsLy9aRnZRMlxuNVU1cWtybVRPc2ExQ2ZWVk5lRGFnMnRmNXA5ZjZvMEl6RkxtOXJSWmRYOWpqcmVualZqeGIwT0Q0blhCU0pNK1xuNGJxdmF4SVRFbjh0OGMxdnZRb2NHNkNvMEtnNXZJTlhIU3J3ZXY4WnhNUmxFejBEdjVpTmlEY0Z1M3g4cnRTRlxuQWwva2dVYlVCT0tyUlM1OVpuOHBIYVpjOWNUdCtDRDBOUGl3eFRScFpEejZYTFFtTmxvRW9oTURkb0RwUGs2alxuZjNRUDRjMTZ6ckNKV2xNTnVLcEw0QnZva2c2bC80Y3JpMmd0S3BWMFVuNDJBMm1EMi9zVkswRlExYWZBVXFaQVxubHBSM3lvMFhBZ01CQUFFQ2dnRUFRbUVlcmlTek80L0o3VHVlc0NIaS9JMjgwbEE1WWdteWVvNGdWR3hoN092aFxua1ZQY0xhZUJmejNNc0gzbmlVU1c0cUFmaE5OME5lZ1Y3OFZzTUZxTnFWLzJ2TXhIV2t6UUVVWmFrcFZCSnlRbVxuOWNWeDkyV2dTdndTOXBFT2RlS1BNUjdzWnVVLzZuMTBzTGJhdHY3d3FPNVBuU2laWTAzYlRPSDBHeStONjFKcVxuWnZ4dzZxUTk5R1Y5dSthbC9yaXRqY1luV1VVSFM2cUxvS29vS21kN2JLVWFCKzJBVkZFdTdkYW16WE5pU0hWd1xuOUU1blRqSUU3WW1JSGh4Z21WdGhid1UzVnRBcXlXclhEOHNmbTNnVGdRaHQvc3ZIMmgzME5xcFFRMHFYdVV0dlxuMkxxcExTTDQzUUh4NUhHUGJsWmRNQ1NsYWdyYkpPMFpKbEp0MERiaVFRS0JnUURnNWRQU1ltV1ZzbTdhVThtRFxuVHlvWVhuZkorYTZFSnJqOFFaZXVGNDAreTM4RVdhbW9FakJYQy9RUkFLaVRaWnVyVEpRSjRKSzdoS3ZjcTUwWlxuMG0wZ05VMHlYSFlaOE10NFlBRnJxNVYyR2pMRmEwdnZNY21HcEpOcExIdW9IRTFQcTAwSzVCR2RMK0M0K01FVFxudm56THJCVndOcjQyTjV1TXVWdE9YK1VqWVFLQmdRREkxU2dxSlFXaWRYNGRJbnpjSzltOTN5U2pxaXZpRkRoblxuVWliV0xNUWc2SWQxMXdRMWFBektYRFdCbDJBdmpSNkRmS0dYU3h5UVloVmdGK0hvMGtQbGlBbjI1dEUxTU9RMVxuVkVJUFFrSG9RTDc2eXVVbmx3OGFlQkI3RVhGdkJZRlRDenBvMzl3SkE3T1RVUU1neFFnRTQ1dzJPRVZDWVJVNVxuNkIvNkdQZjdkd0tCZ0VBYkl3TVFTemFka3gybXpvVWdkclpWanozQ2J5MTZRRGFmbDlYbkhycnNsTHN2dDcyUlxuMmJlbVQ1N2RHakJSd1VyVjNFam9lOHI1NldkUWY4cVJnY3V4aGR5NklHd0tpR1U2c0l6NWJ2UW9TWnFlbUJtSlxuanBGVHJqVkhtY1AwdFZEbFdtT2xXU1ArTnNkakdONUE2Ui9CWUtjSTQ3VmVxbmMxaXNKUGNXaGhBb0dBU2xZclxuOGIyV1lsWGZBbVc3bGJ0ZGMxeGZNWDFSbFBNckZZTmhBTEc5UHZrUyt0bEZvNGNLQlBVd2tQRTVGeUREQktSalxuSTU4WHlaR2ttOWI3TmJSdWtVMVRjdUpvMmFscEx4d3EvMzdrNmxUYzIvT3g4bitGaGFTZlpRNUYwSFJYTkNmWlxuek92ZmhDNnNqNERUQ1pRa3JicWFIYStpWXptUUk4ajUzUmJVNDIwQ2dZQWVtQTM4dllDNUhPcDZZMlRDaC9TR1xuZXFGZklSZXd0L2pITnRhNVRyUE9tanhYZThNdzl6Wm5qMEFxUHdQU3hER2hzQnlzNzB1SlYxdS9Kd2VheHFmVlxuZUF4STEreDZ6V0M4Uk1IcDY1T1N6R3lIK0lxK3hiTUJIdVQ3eGZVNFdWZGdSV2dxdGNUUzVjZzBMQ3R2T0xhMVxuaG0zRXh3TkhFMnBlSUFDUDNkMGt3dz09XG4tLS0tLUVORCBQUklWQVRFIEtFWS0tLS0tXG4iLAogICJjbGllbnRfZW1haWwiOiAidGVzdC04NjdAc3JjZC10ZXN0LmlhbS5nc2VydmljZWFjY291bnQuY29tIiwKICAiY2xpZW50X2lkIjogIjExNjA5OTQ3NDMwMzQwMzk5ODExMSIsCiAgImF1dGhfdXJpIjogImh0dHBzOi8vYWNjb3VudHMuZ29vZ2xlLmNvbS9vL29hdXRoMi9hdXRoIiwKICAidG9rZW5fdXJpIjogImh0dHBzOi8vYWNjb3VudHMuZ29vZ2xlLmNvbS9vL29hdXRoMi90b2tlbiIsCiAgImF1dGhfcHJvdmlkZXJfeDUwOV9jZXJ0X3VybCI6ICJodHRwczovL3d3dy5nb29nbGVhcGlzLmNvbS9vYXV0aDIvdjEvY2VydHMiLAogICJjbGllbnRfeDUwOV9jZXJ0X3VybCI6ICJodHRwczovL3d3dy5nb29nbGVhcGlzLmNvbS9yb2JvdC92MS9tZXRhZGF0YS94NTA5L3Rlc3QtODY3JTQwc3JjZC10ZXN0LmlhbS5nc2VydmljZWFjY291bnQuY29tIgp9Cg=="
	gcsConfig        = GCSConfig{
		Bucket: "domain-test-fs",
		Scope:  storage.ScopeReadWrite,
	}

	gcsRootFixture     = "foo"
	gcsOpenFixture     = "open.txt"
	gcsCreateFixture   = "create.txt"
	gcsContentsFixture = []byte("42 foobar 42 fooqux 42 foo 42 bar 42 qux 42 foobarqux 42 42 42\n")
)

type GoogleCloudStorageSuite struct {
	client *GCSClient
}

func (g *GoogleCloudStorageSuite) SetUpSuite(c *C) {
	var err error
	gcsConfig.Key, err = base64.StdEncoding.DecodeString(gcsConfigFixture)
	c.Assert(err, IsNil, Commentf("failed base64 decode of Google Cloud Storage key fixture"))

	g.client, err = NewGCSClient(gcsRootFixture, gcsConfig)
	c.Assert(g.client, NotNil)
	c.Assert(err, IsNil)

	f, err := g.client.Create(gcsOpenFixture)
	c.Assert(err, IsNil)
	n, err := f.Write(gcsContentsFixture)
	c.Assert(err, IsNil)
	c.Assert(n, Equals, len(gcsContentsFixture))
	c.Assert(f.Close(), IsNil)
}

func (g *GoogleCloudStorageSuite) TearDownSuite(c *C) {
	c.Assert(g.client.Remove(gcsOpenFixture), IsNil)
	c.Assert(g.client.Close(), IsNil)
}

func (g *GoogleCloudStorageSuite) Test_GCS_Client_Open(c *C) {
	_, err := g.client.Open("404.txt")
	c.Assert(err, NotNil)

	f, err := g.client.Open(gcsOpenFixture)
	c.Assert(err, IsNil)
	c.Assert(f.GetFilename(), Equals, path.Join(gcsRootFixture, gcsOpenFixture))
	contents, err := ioutil.ReadAll(f)
	c.Assert(err, IsNil)
	c.Assert(contents, DeepEquals, gcsContentsFixture)
	c.Assert(f.Close(), IsNil)
}

func (g *GoogleCloudStorageSuite) Test_GCS_Client_Create(c *C) {
	f1, err := g.client.Create(gcsCreateFixture)
	c.Assert(err, IsNil)
	_, err = f1.Write(gcsContentsFixture)
	c.Assert(err, IsNil)
	c.Assert(f1.Close(), IsNil)

	f2, err := g.client.Open(gcsCreateFixture)
	c.Assert(err, IsNil)
	contents, err := ioutil.ReadAll(f2)
	c.Assert(err, IsNil)
	c.Assert(contents, DeepEquals, gcsContentsFixture)
	c.Assert(f2.Close(), IsNil)

	c.Assert(g.client.Remove(gcsCreateFixture), IsNil)
}

func (g *GoogleCloudStorageSuite) Test_GCS_Client_Stat(c *C) {
	fi, err := g.client.Stat(gcsOpenFixture)
	c.Assert(err, IsNil)

	c.Assert(fi.Name(), Equals, path.Join(gcsRootFixture, gcsOpenFixture))
	// Google Cloud Storage returns 0 no matter what
	c.Assert(fi.Size(), Equals, int64(len(gcsContentsFixture)))
	c.Assert(fi.Mode(), Equals, os.FileMode(0))
	c.Assert(fi.ModTime().IsZero(), Equals, false)
	c.Assert(fi.IsDir(), Equals, false)
	c.Assert(fi.Sys(), Equals, nil)
}

func (g *GoogleCloudStorageSuite) Test_GCS_Client_Read(c *C) {
	filename := path.Join(gcsRootFixture, "read.txt")

	// Create file to read
	{
		f, err := g.client.Create(filename)
		c.Assert(err, IsNil)

		n, err := f.Write(gcsContentsFixture)
		c.Assert(err, IsNil)
		c.Assert(n, Equals, len(gcsContentsFixture))

		c.Assert(f.Close(), IsNil)
	}

	f, err := g.client.Open(filename)
	c.Assert(err, IsNil)

	buf := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	n, err := f.Read(buf)
	c.Assert(err, IsNil)
	c.Assert(n, Equals, len(buf))
	c.Assert(buf, DeepEquals, gcsContentsFixture[:10])

	c.Assert(f.Close(), IsNil)
	c.Assert(g.client.Remove(filename), IsNil)
}
