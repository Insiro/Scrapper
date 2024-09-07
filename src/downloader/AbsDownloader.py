from abc import ABCMeta, abstractmethod
class AbsDownloader(metaclass = ABCMeta):

    @abstractmethod
    def scrap(self,url:str):
        pass
