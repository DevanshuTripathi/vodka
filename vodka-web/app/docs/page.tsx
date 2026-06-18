import { fetchREADME, fetchREADMESections } from "../lib/github";
import MarkdownRenderer from "../components/MarkdownRenderer";
import Navbar from "../components/Navbar";
import FeatureCards from "../components/FeatureCards";
import WhyVodkaCards from "../components/WhyVodkaCards";
import QuickStartSection from "../components/QuickStartSection";
import InstallationSection from "../components/InstallationSection";
import ProjectScaffoldingSection from "../components/ProjectScaffoldingSection";
import CoreConceptsSection from "../components/CoreConceptsSection";
import MiddlewareSection from "../components/MiddlewareSection";
import ValidationSection from "../components/ValidationSection";
import AuthenticationSection from "../components/AuthenticationSection";
import TemplatesSection from "../components/TemplatesSection";
import SPASupportSection from "../components/SPASupportSection";
import RoadmapSection from "../components/RoadmapSection";
import DemoSection from "../components/DemoSection";
import MinimalAPISection from "../components/MinimalAPISection";
import RequestIDSection from "../components/RequestIDSection";
import TemplateRenderingSection from "../components/TemplateRenderingSection";
import AdditionalInfoSection from "../components/AdditionalInfoSection";


export const revalidate = 3600;

export default async function DocsPage() {
  const { installationContent, remainingContent } = await fetchREADMESections();

  return (
    <div>
      <Navbar />
      <div className="bg-slate-100 min-h-screen text-slate-900 pt-0">
        {" "}
        <div className="max-w-6xl mx-auto flex items-start">
          {" "}
          <aside className="hidden lg:block w-80 shrink-0 sticky top-16 h-[calc(100vh-4rem)] border-r border-slate-300 bg-slate-50 p-10 overflow-y-auto">
            {" "}
            <h3 className="font-bold text-xl mb-6">Documentation</h3>
            <nav className="space-y-4 text-base text-slate-700">
              {" "}
              <a
                href="#features"
                className="block px-3 py-0.2 rounded-lg hover:bg-slate-300 hover:text-black transition-all duration-200"
              >
                Features
              </a>
              <a
                href="#why-vodka"
                className="block px-3 py-0.2 rounded-lg hover:bg-slate-300 hover:text-black transition-all duration-200"
              >
                Why Vodka
              </a>
              <a
                href="#installation"
                className="block px-3 py-0.2 rounded-lg hover:bg-slate-300 hover:text-black transition-all duration-200"
              >
                Installation
              </a>
              <a
                href="#quick-start"
                className="block px-3 py-0.2 rounded-lg hover:bg-slate-300 hover:text-black transition-all duration-200"
              >
                Quick Start
              </a>
              <a
                href="#demo"
                className="block px-3 py-0.2 rounded-lg hover:bg-slate-300 hover:text-black transition-all duration-200"
              >
                Demo
              </a>
              <a
                href="#project-scaffolding"
                className="block px-3 py-0.2 rounded-lg hover:bg-slate-300 hover:text-black transition-all duration-200"
              >
                Project Scaffolding
              </a>
              <a
                href="#minimal-api"
                className="block px-3 py-0.2 rounded-lg hover:bg-slate-300 hover:text-black transition-all duration-200"
              >
                Minimal API
              </a>
              <a
                href="#core-concepts"
                className="block px-3 py-0.2 rounded-lg hover:bg-slate-300 hover:text-black transition-all duration-200"
              >
                Core Concepts
              </a>
              <a
                href="#middleware"
                className="block px-3 py-0.2 rounded-lg hover:bg-slate-300 hover:text-black transition-all duration-200"
              >
                Middleware
              </a>
              <a
                href="#request-id"
                className="block px-3 py-0.2 rounded-lg hover:bg-slate-300 hover:text-black transition-all duration-200"
              >
                Request ID
              </a>
              <a
                href="#validation"
                className="block px-3 py-0.2 rounded-lg hover:bg-slate-300 hover:text-black transition-all duration-200"
              >
                Validation
              </a>
              <a
                href="#authentication"
                className="block px-3 py-0.2 rounded-lg hover:bg-slate-300 hover:text-black transition-all duration-200"
              >
                Authentication
              </a>
              <a
                href="#templates"
                className="block px-3 py-0.2 rounded-lg hover:bg-slate-300 hover:text-black transition-all duration-200"
              >
                Templates
              </a>
              <a
                href="#template-rendering"
                className="block px-3 py-0.2 rounded-lg hover:bg-slate-300 hover:text-black transition-all duration-200"
              >
                Template Rendering
              </a>
              <a
                href="#spa-support"
                className="block px-3 py-0.2 rounded-lg hover:bg-slate-300 hover:text-black transition-all duration-200"
              >
                SPA Support
              </a>
              <a
                href="#additional-info"
                className="block px-3 py-0.2 rounded-lg hover:bg-slate-300 hover:text-black transition-all duration-200"
              >
                Additional Info
              </a>
              <a
                href="#roadmap"
                className="block px-3 py-0.2 rounded-lg hover:bg-slate-300 hover:text-black transition-all duration-200"
              >
                Roadmap
              </a>
            </nav>
          </aside>
          <main className="flex-1 px-8 py-4">
            <div className="mb-8">
              <p className="uppercase tracking-[0.25em] text-slate-500 text-xs mb-3">
                Documentation
              </p>

              <h1 className="text-3xl font-bold mb-3">Vodka Documentation</h1>

              <p className="text-slate-600 text-sm max-w-2xl leading-6">
                Learn how to build modern Go applications with Vodka using
                guides, examples, middleware documentation, and full-stack
                workflows.
              </p>
            </div>

            <div
              id="features"
              className="border border-slate-200 rounded-xl p-8 my-8 bg-white"
            >
              <h2 className="text-2xl font-bold mb-4">Features</h2>
              <FeatureCards />
            </div>

            <div
              id="why-vodka"
              className="border border-slate-200 rounded-xl p-8 my-8 bg-white"
            >
              <h2 className="text-2xl font-bold mb-4">Why Vodka?</h2>

              <p className="text-slate-600 mb-6">
                Vodka combines fast backend iteration with a modern full-stack
                workflow without the configuration overhead of heavier
                frameworks.
              </p>

              <WhyVodkaCards />
            </div>

            <InstallationSection />
            <QuickStartSection />
            <DemoSection />
            <ProjectScaffoldingSection />
            <MinimalAPISection />
            <CoreConceptsSection />
            <MiddlewareSection />
            <RequestIDSection />
            <ValidationSection />
            <AuthenticationSection />
            <TemplatesSection />
            <TemplateRenderingSection />
            <SPASupportSection />
            <AdditionalInfoSection />
            <RoadmapSection />
          </main>
        </div>
      </div>
    </div>
  );
}
