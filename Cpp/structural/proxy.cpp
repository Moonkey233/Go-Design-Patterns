#include "proxy.h"
namespace proxy {
    std::pair<int, std::string> Application::request(const std::string url, const std::string method) const{
        if (url == "/proxy" && method == "GET") return std::make_pair(200, "OK");
        return std::make_pair(404, "Not Found");
    }
    Nginx::Nginx(Application *app) { m_app = app; }
    bool Nginx::checkAccess(const std::string& url, const std::string& method) {
        if (url.empty() || (method != "GET" && method != "POST")) return false;
        return true;
    }
    std::pair<int, std::string> Nginx::request(const std::string url, const std::string method) const{
        if (checkAccess(url, method)) return m_app->request(url, method);
        return std::make_pair(403, "Forbidden");
    }
    void TestProxy() {
        Application app{};
        const Nginx nginx(&app);
        std::vector<std::pair<std::string, std::string>> requests = {
            {"", "GET"},
            {"/proxy", "POST"},
            {"/proxy", "GET"},
          };
        for (auto &req : requests) {
            std::cout << "nginx: checking access with url = " << req.first << ", method = " << req.second << std::endl;
            auto res = nginx.request(req.first, req.second);
            std::cout << "nginx: got response: " << res.first << ", " << res.second << std::endl;
        }
    }
}