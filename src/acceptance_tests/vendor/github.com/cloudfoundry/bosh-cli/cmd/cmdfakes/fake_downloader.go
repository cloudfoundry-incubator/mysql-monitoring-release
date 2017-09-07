// Code generated by counterfeiter. DO NOT EDIT.
package cmdfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/cmd"
)

type FakeDownloader struct {
	DownloadStub        func(blobstoreID, sha1, prefix, dstDirPath string) error
	downloadMutex       sync.RWMutex
	downloadArgsForCall []struct {
		blobstoreID string
		sha1        string
		prefix      string
		dstDirPath  string
	}
	downloadReturns struct {
		result1 error
	}
	downloadReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDownloader) Download(blobstoreID string, sha1 string, prefix string, dstDirPath string) error {
	fake.downloadMutex.Lock()
	ret, specificReturn := fake.downloadReturnsOnCall[len(fake.downloadArgsForCall)]
	fake.downloadArgsForCall = append(fake.downloadArgsForCall, struct {
		blobstoreID string
		sha1        string
		prefix      string
		dstDirPath  string
	}{blobstoreID, sha1, prefix, dstDirPath})
	fake.recordInvocation("Download", []interface{}{blobstoreID, sha1, prefix, dstDirPath})
	fake.downloadMutex.Unlock()
	if fake.DownloadStub != nil {
		return fake.DownloadStub(blobstoreID, sha1, prefix, dstDirPath)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.downloadReturns.result1
}

func (fake *FakeDownloader) DownloadCallCount() int {
	fake.downloadMutex.RLock()
	defer fake.downloadMutex.RUnlock()
	return len(fake.downloadArgsForCall)
}

func (fake *FakeDownloader) DownloadArgsForCall(i int) (string, string, string, string) {
	fake.downloadMutex.RLock()
	defer fake.downloadMutex.RUnlock()
	return fake.downloadArgsForCall[i].blobstoreID, fake.downloadArgsForCall[i].sha1, fake.downloadArgsForCall[i].prefix, fake.downloadArgsForCall[i].dstDirPath
}

func (fake *FakeDownloader) DownloadReturns(result1 error) {
	fake.DownloadStub = nil
	fake.downloadReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDownloader) DownloadReturnsOnCall(i int, result1 error) {
	fake.DownloadStub = nil
	if fake.downloadReturnsOnCall == nil {
		fake.downloadReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.downloadReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDownloader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.downloadMutex.RLock()
	defer fake.downloadMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDownloader) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ cmd.Downloader = new(FakeDownloader)
