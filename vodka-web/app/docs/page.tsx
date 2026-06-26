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
          <aside className="hidden lg:block w-64 shrink-0 sticky top-20 h-[calc(100vh-5rem)] bg-white p-4">
            <h3 className="text-sm font-medium text-black text-slate-700 tracking-[0.12em] uppercase mb-3 px-2">
              Documentation
            </h3>
            <nav className="flex flex-col text-xs">
              <div className="h-px bg-slate-100 my-1.5 mx-2" />

              <a
                href="#features"
                className="px-3 py-1.5 text-slate-700 border-l-2 border-transparent hover:text-black hover:bg-slate-100 transition-all duration-150"
              >
                Features
              </a>
              <a
                href="#why-vodka"
                className="px-3 py-1.5 text-slate-700 border-l-2 border-transparent hover:text-black hover:bg-slate-100 transition-all duration-150"
              >
                Why Vodka
              </a>

              <div className="h-px bg-slate-100 my-1.5 mx-2" />

              <a
                href="#installation"
                className="px-3 py-1.5 text-slate-700 border-l-2 border-transparent hover:text-black hover:bg-slate-100 transition-all duration-150"
              >
                Installation
              </a>
              <a
                href="#quick-start"
                className="px-3 py-1.5 text-slate-700 border-l-2 border-transparent hover:text-black hover:bg-slate-100 transition-all duration-150"
              >
                Quick Start
              </a>
              <a
                href="#demo"
                className="px-3 py-1.5 text-slate-700 border-l-2 border-transparent hover:text-black hover:bg-slate-100 transition-all duration-150"
              >
                Demo
              </a>
              <a
                href="#project-scaffolding"
                className="px-3 py-1.5 text-slate-700 border-l-2 border-transparent hover:text-black hover:bg-slate-100 transition-all duration-150"
              >
                Project Scaffolding
              </a>
              <a
                href="#minimal-api"
                className="px-3 py-1.5 text-slate-700 border-l-2 border-transparent hover:text-black hover:bg-slate-100 transition-all duration-150"
              >
                Minimal API
              </a>

              <div className="h-px bg-slate-100 my-1.5 mx-2" />

              <a
                href="#core-concepts"
                className="px-3 py-1.5 text-slate-700 border-l-2 border-transparent hover:text-black hover:bg-slate-100 transition-all duration-150"
              >
                Core Concepts
              </a>
              <a
                href="#middleware"
                className="px-3 py-1.5 text-slate-700 border-l-2 border-transparent hover:text-black hover:bg-slate-100 transition-all duration-150"
              >
                Middleware
              </a>
              <a
                href="#request-id"
                className="px-3 py-1.5 text-slate-700 border-l-2 border-transparent hover:text-black hover:bg-slate-100 transition-all duration-150"
              >
                Request ID
              </a>
              <a
                href="#validation"
                className="px-3 py-1.5 text-slate-700 border-l-2 border-transparent hover:text-black hover:bg-slate-100 transition-all duration-150"
              >
                Validation
              </a>
              <a
                href="#authentication"
                className="px-3 py-1.5 text-slate-700 border-l-2 border-transparent hover:text-black hover:bg-slate-100 transition-all duration-150"
              >
                Authentication
              </a>

              <div className="h-px bg-slate-100 my-1.5 mx-2" />

              <a
                href="#templates"
                className="px-3 py-1.5 text-slate-700 border-l-2 border-transparent hover:text-black hover:bg-slate-100 transition-all duration-150"
              >
                Templates
              </a>
              <a
                href="#template-rendering"
                className="px-3 py-1.5 text-slate-700 border-l-2 border-transparent hover:text-black hover:bg-slate-100 transition-all duration-150"
              >
                Template Rendering
              </a>
              <a
                href="#spa-support"
                className="px-3 py-1.5 text-slate-700 border-l-2 border-transparent hover:text-black hover:bg-slate-100 transition-all duration-150"
              >
                SPA Support
              </a>
              <a
                href="#additional-info"
                className="px-3 py-1.5 text-slate-700 border-l-2 border-transparent hover:text-black hover:bg-slate-100 transition-all duration-150"
              >
                Additional Info
              </a>
              <a
                href="#roadmap"
                className="px-3 py-1.5 text-slate-700 border-l-2 border-transparent hover:text-black hover:bg-slate-100 transition-all duration-150"
              >
                Roadmap
              </a>
            </nav>
          </aside>
          <main className="flex-1 px-8 py-4">
            <div className="mb-8">
              <p className="uppercase tracking-[0.25em] text-slate-500 text-lg mb-3">
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
